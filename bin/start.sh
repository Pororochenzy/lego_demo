./stup.sh start mainte mainte ./conf/mainte.yaml
sleep 1
./stup.sh start worker_1 worker ./conf/worker_1.yaml 
sleep 1
./stup.sh start gateway_1 gateway ./conf/gateway_1.yaml 
