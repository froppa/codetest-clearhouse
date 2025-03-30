#!/bin/bash
set -u

BASE_URL="http://localhost:8081"
COMPANY_NAME="Clearhouse"
COUNTRY="Denmark"
EMAIL="contact@acme.com"
OWNER_NAME="John Doe"
SSN=$(shuf -i 1000000000-9999999999 -n 1)
SSN_FORMATTED=$(echo "$SSN" | sed 's/\([0-9]\{6\}\)\([0-9]\{4\}\)/\1-\2/')

PASSED_TEST_COUNT=0
FAILED_TEST_COUNT=0

test_step() {
  local label="$1"
  local result="$2"
  local expected="${3:-}"

  result="$(echo "$result" | awk '{$1=$1; print}')"

  echo "=> TEST: $label"
  if echo "$result" | jq -e . >/dev/null 2>&1; then
    echo "== Result: $(echo "$result" | jq -c .)"
    ((PASSED_TEST_COUNT++))
  else
    echo "== Result: $result"
    if [[ -n "$expected" && ("$result" != "$expected" || "$result" == "") ]]; then
      ((FAILED_TEST_COUNT++))
    else
      ((PASSED_TEST_COUNT++))
    fi
  fi
  echo
}

# Create Company
CREATE_COMPANY=$(curl -s -X POST "$BASE_URL/companies" \
  -H "Content-Type: application/json" \
  -d "{\"name\": \"$COMPANY_NAME\", \"country\": \"$COUNTRY\", \"email\": \"$EMAIL\"}")

COMPANY_ID=$(echo "$CREATE_COMPANY" | jq -r ".id")

# List Companies
LIST_COMPANIES=$(curl -s -X GET "$BASE_URL/companies")

# Get Company Details
GET_COMPANY=$(curl -s -X GET "$BASE_URL/companies/$COMPANY_ID")

# Update Company
UPDATE_COMPANY=$(curl -s -X PUT "$BASE_URL/companies/$COMPANY_ID" \
  -H "Content-Type: application/json" \
  -d "{\"name\": \"Updated $COMPANY_NAME\"}")

# Add Owner
ADD_OWNER=$(curl -s -X POST "$BASE_URL/companies/$COMPANY_ID/owners" \
  -H "Content-Type: application/json" \
  -d "{\"name\": \"$OWNER_NAME\", \"ssn\": \"$SSN_FORMATTED\"}")

# Check SSN with Admin Access
CHECK_SSN_ADMIN=$(curl -s -X POST "$BASE_URL/auth/verify?group=admin" \
  -H "Content-Type: application/json" \
  -d "{\"ssn\": \"$SSN_FORMATTED\"}")

# Check SSN with Restricted Access
CHECK_SSN_RESTRICTED=$(curl -s -X POST "$BASE_URL/auth/verify?group=restricted" \
  -H "Content-Type: application/json" \
  -d "{\"ssn\": \"$SSN\"}")

# Run Tests
echo "==========================="
echo "=> Testing API Endpoints <="
echo "==========================="

test_step "Create Company" "$CREATE_COMPANY"
test_step "List Companies" "$LIST_COMPANIES"
test_step "Get Company" "$GET_COMPANY"
# test_step "Update Company" "$UPDATE_COMPANY"
test_step "Add Owner" "$ADD_OWNER"
test_step "Check SSN (Admin)" "$CHECK_SSN_ADMIN"
test_step "Check SSN (Restricted)" "$CHECK_SSN_RESTRICTED" "Invalid SSN Format"

echo "================================================="
echo "=> Test Summary"
echo "================================================="
echo "✅ Passed: $PASSED_TEST_COUNT"
echo "❌ Failed: $FAILED_TEST_COUNT"
echo
