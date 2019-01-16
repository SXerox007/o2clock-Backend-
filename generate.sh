# Curl Register User
curl -X POST -k http://localhost:5051/v1/user/register/78787 -d '{"email":"sumit@hotmail.com"}'

# Curl Access Token
curl -X GET http://localhost:5051/v1/user/accesstoken

# Curl Logout User
curl -X POST http://localhost:5051/v1/user/logout -d '{"access_token":""}'

#Curl Login User
curl -X POST http://localhost:5051/v1/user/login -d '{"username_email":"sumitthakur769@gmail.com","password":"Xerox007#"}'

# Curl multipart Home Verify User
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -v -F file="/Users/sumitthakur/Downloads/hash.jpg" -X POST http://localhost:5051/v1/user/verify

## Here start the chat curl
# Curl Get all the users
curl -X GET http://localhost:5051/v1/chat/userslist -d '{"access_token":"eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDYyNjE4MjUsInN1YiI6Ik9iamVjdElEKFwiNWMyYTE1MDVmYTAyNmU5Y2U0NzA5ZWJiXCIpIiwidXNlcm5hbWUiOiIifQ.5K4IH2DWzGlaH5nLAc6RCmNNcB0yZtmDxishzLEQgEfPU3PdDX0z61SxP-vQet8UM_530fhuNi-db_sopqsq7quQe0uUCmjm5Cyd5ei9GtK32-Jilzx_4nm86_UAC0Q11DugL1DYLKumMiJOKyCNl9JgDgWNVEVsRh5fs3NWI0DCizuGq2dCwbG3-U0lnoiMuwBpY67G5rXtF4T-PEfrDxXGelFXMwmTWMDRpDVY0OzhK8MPgCPtNSV8KW9gt-UDTWAgllDnALUN0sv5UafTFqh4b5dX9Gw04cm1gFDN1WfFGBpCfU3DHlUe4qJr_zQcbm8y_RcJwVqwj4JAingZ4g"}'

# source keys
source secure-keys/keys

# secure key
openssl genrsa -out secure-keys/o2clock.rsa