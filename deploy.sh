
# convenience script to deploy the SAM stack
# TODO: replae with pipeline

#aws cloudformation delete-stack --stack-name lambda-go-structured-inputs

aws cloudformation deploy \
    --template-file structured.sam.yaml \
    --stack-name lambda-go-structured-inputs \
    --capabilities CAPABILITY_IAM

