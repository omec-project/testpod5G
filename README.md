# testpod5G

This is a developed tool to simulate 5G NFs like AMF, SMF, NRF, PCF, UDM. The real NF to test, can be like SMF while all other NFs shall be simulated using testpod.
The Testpod can be run locally on any environment having miniKube via helm test package file present in repository OR locally as in binary mode.


On minikube
Example- to test SMF(precondition- UPF should be running)
helm install smftest ./helm/smf/


In Binary mode
SMF
./smf -smfcfg ../../config/smfcfg.yaml -uerouting ../../config/uerouting.yaml

UPF
./pfcpiface -config ../conf/upf.json

TestPod App
./testpod amf
