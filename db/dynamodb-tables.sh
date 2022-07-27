aws dynamodb create-table \
    --table-name Truck \
    --attribute-definitions \
        AttributeName=UUID,AttributeType=S \
    --key-schema \
        AttributeName=UUID,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
    --endpoint-url http://localhost:8000

aws dynamodb create-table \
    --table-name Warehouse \
    --attribute-definitions \
        AttributeName=UUID,AttributeType=S \
    --key-schema \
        AttributeName=UUID,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
    --endpoint-url http://localhost:8000

aws dynamodb create-table \
    --table-name Vineyard \
    --attribute-definitions \
        AttributeName=UUID,AttributeType=S \
    --key-schema \
        AttributeName=UUID,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
    --endpoint-url http://localhost:8000

aws dynamodb create-table \
    --table-name GrapeType \
    --attribute-definitions \
        AttributeName=UUID,AttributeType=S \
    --key-schema \
        AttributeName=UUID,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
    --endpoint-url http://localhost:8000

aws dynamodb create-table \
    --table-name Reception \
    --attribute-definitions \
        AttributeName=UUID,AttributeType=S \
    --key-schema \
        AttributeName=UUID,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
    --endpoint-url http://localhost:8000

aws dynamodb create-table \
    --table-name Maceration \
    --attribute-definitions \
        AttributeName=UUID,AttributeType=S \
    --key-schema \
        AttributeName=UUID,KeyType=HASH \
    --provisioned-throughput \
        ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --table-class STANDARD \
    --endpoint-url http://localhost:8000

aws dynamodb put-item \
    --table-name Truck  \
    --item \
        '{"UUID": {"S": "a57ccb31-a58b-4f6e-817e-b0a85fc652e4"}, "License": {"S": "EEOO-990"}}' \
    --endpoint-url http://localhost:8000

aws dynamodb put-item \
    --table-name Vineyard  \
    --item \
        '{"UUID": {"S": "5da1acf8-d613-4af0-8b82-15493346bf56"}, "Name": {"S": "Los Andes"}}' \
    --endpoint-url http://localhost:8000

aws dynamodb put-item \
    --table-name GrapeType  \
    --item \
        '{"UUID": {"S": "6555c14c-07da-44c0-a85e-094c250448ea"}, "Name": {"S": "Rosada"}}' \
    --endpoint-url http://localhost:8000

aws dynamodb put-item \
    --table-name Warehouse  \
    --item \
        '{"UUID": {"S": "93f24e5f-59b8-488d-a583-dc6d948140bb"}, "Name": {"S": "EEEEEE"}, "IsEmpty": {"BOOL": true}}' \
    --endpoint-url http://localhost:8000


