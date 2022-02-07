<!--
Copyright 2021-present Open Networking Foundation
SPDX-License-Identifier: Apache-2.0

-->

# testpod5G

Testpod is a developer tool to fast track unit testing of any NF by simulating other NFs like AMF, PCF, UDM, NRF and UPF(only PFCP functionality) for SMF unit testing.

## Supported Features
1. Testpod supports  SMF unit testing. 
2. The User can run an actual SMF application and testpod application to simulate other NFs which SMF interacts with. 
3. The User can control behaviour of other NFs which reside in Testpod to verify SMF handling, like negative responses, timeouts, invalid response, etc from peer network functions. 

## TestPod Setup
![TestPod](/docs/images/TestPod.png)

## Planned Features

1. PCC Rule support from PCF(Testpod) to SMF
2. SMPolicy Notify Callback from PCF to SMF.
3. Run a set of predefined test cases against SM
4. Extend Functionality to test other NF functions 


More details of the testPod can be found at - https://docs.sd-core.opennetworking.org/master/developer/testpod.html 
