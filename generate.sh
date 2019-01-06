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

# source keys
source secure-keys/keys

# secure key
openssl genrsa -out secure-keys/o2clock.rsa