# Curl Register User
curl -X POST -k http://localhost:5051/v1/user/register/78787 -d '{"email":"sumit@hotmail.com"}'

# Curl Access Token
curl -X GET http://localhost:5051/v1/user/accesstoken

# Curl Logout User
curl -X POST http://localhost:5051/v1/user/logout -d '{"access_token":""}'

#Curl Login User
curl -X POST http://localhost:5051/v1/user/login -d '{"username_email":"sumit@bonfleet.com","password":"12345678"}'

# Curl multipart Home Verify User
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -v -F file="/Users/sumitthakur/Downloads/hash.jpg" -X POST http://localhost:5051/v1/user/verify

## Here start the chat curl
# Curl Get all the users
curl -X GET http://localhost:5051/v1/chat/userslist

# Curl forgot password
curl -i -H "Content-Type: application/json" -X GET 'http://localhost:5051/v1/user/password/forgot?email=sumit@bonfleet.com'
# OR
curl -i -H "Content-Type: application/json" -X GET 'https://92de1297.ngrok.io/v1/user/password/forgot?email=sumit@bonfleet.com'

# Curl reset password
curl -i -H "Content-Type: application/json" -X GET 'https://92de1297.ngrok.io/v1/user/setpassword/xyx'


# source keys
source secure-keys/keys

# secure key
openssl genrsa -out secure-keys/o2clock.rsa

# mongodb clear collection
db.all_single_chats.remove({})
