#!/bin/bash
# for i in {1..50}; do go test || break; done > test_results
# go test -list . | grep -v ok | tr '\n' ' '

# TestBasic TestMove TestLimp TestConcurrent TestConcurrentUnreliable
for t in TestBasic TestMove TestLimp TestConcurrent TestConcurrentUnreliable
do
  echo $t
  count=0
  n=100
  for i in $(seq 1 $n)
  do
    go test -run "^${t}$" -timeout 2m > ./log-${t}-${i}.log
    result=$(grep -E '^PASS$' log-${t}-${i}.log| wc -l)
    count=$((count + result))
    if [ $result -eq 1 ]; then
       rm ./log-${t}-${i}.log
    fi
  done
  echo "$count/$n"
done