# Gomariadb service app 
golang web app which use gorm as ORM and mariadb as underlying database. Publishes messages to kafka topic whenever there is an event.

## endpoints : 
        http://localhost:8080/api/v1/register
        {
            "firstName" : "Bharath",
            "lastName" : "Kuppala",
            "email" : "b@gmail.com",
            "bloodGroup" : "B+",
            "password" : "pass",
            "userName": "Bharath Kuppala"
        }

        http://localhost:8080/api/v1/login
        {
             "email" : "b@gmail.com",
             "password" : "pass"
        }

## Just for reference
DSN: bharath:password@tcp(localhost:3306)/mariadb_test?checkConnLiveness=false&parseTime=true&maxAllowedPacket=0

## kafka commands (To run locally)
        1. JMX_PORT=8004 bin/kafka-server-start.sh config/server.properties ---> to start kafka
        2. bin/zookeeper-server-start.sh config/zookeeper.properties ---> to start zookeeper
        3. bin/cmak -Dconfig.file=conf/application.conf -Dhttp.port=9000 ----> to start kafka manager(gui)

## Commands needs to be executed to run services inside container
        1. docker-compose up -d ----> for background mode
        2. docker-compose up ----> for foreground mode

## To get inside container
        1. docker exec -it kafka bash or sh
        2. navigate to /opt/kafka_2.12-2.3.0/bin/
        3. kafka-topics.sh --list --zookeeper zookeeper:2181 ---> list of topics available 

## Output:
    __consumer_offsets
     mariadb-events

## Command to see the messages which are published to topic
        kafka-console-consumer.sh --bootstrap-server kafka:9092 --topic mariadb-events --from-beginning
        Bharath Kuppala student successfully registered
        Bharath Kuppala successfully logged in at: 2020-09-21 04:31:57
        Bharath Kuppala successfully logged in at: 2020-09-21 05:21:40
       


