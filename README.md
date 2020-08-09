# jwt-user

Good References

https://github.com/victorsteven/jwt-best-practices

https://github.com/ishrivatsa/user

https://github.com/open-policy-agent/opa

https://github.com/open-policy-agent/example-api-authz-go

https://livebook.manning.com/book/microservices-security-in-action/welcome/v-7/

https://medium.com/@cgrant/golang-oauth2-google-example-for-web-and-api-59187ce8b119

Introduction

Security is a critical aspect developing a web application. Auth-N and Auth-Z are 2 key parts of it. 
In this module. we will add auth-N and auth-z support for tiny url

User Management

There are 2 ways of creating the user. 
- He can go ahead and provide all the details
- He can use google, facebook etc API for managing the user account. The app does not need to store any password

JWT

Once the user logs in, we need to ensure the communication between the user and the application is secure
JWT mechanism is used to secure it in this example. The app generates access and refresh token. 
The user inputs access token with bearer token. The application retrieves the token, identifies the user and performs the operation
Since there is a timeout involved with access token, the user can use refresh token to get a new access token
When the user logs out, the token is destroyed

Authorization

The applications provides different roles. E.g. the paid customers are only able to generate custom url vs free users.
We use a policy engine to store the policy details. i.e User X has permission to run attribute Y on Resource Z
We have used OPA which is a cloud native Policy Enginer written in go.

Setting Up

1. Setting Google Account for OAuth

a. Go to https://console.developers.google.com
b. Login and Create a New Project
c. In the sidebar under "APIs & Services", select Credentials
d. In the Credentials tab, select the Create credentials drop-down list, and choose OAuth client ID
e. Under Application type, select Web application
f. In Authorized redirect URI use http://localhost:8080/v1/auth (more details later)
g. Press the Create button and copy the generated client ID and client secret into a following file

E.g 

samdeep-a02:jwt-user samdeep$ cat creds.json 
{
  "cid":"PLS Replace it with what your ID",
  "csecret":"Pls replace this with your Secret"
}
samdeep-a02:jwt-user samdeep$ 

2. Install Redis

Redis is used for storing the token
On Mac, you can use brew to install it

To run redis run, "redis-server /usr/local/etc/redis.conf"

3. Install OPA

curl -L -o opa https://openpolicyagent.org/downloads/latest/opa_darwin_amd64
chmod 755 ./opa
./opa run

You can also download and run OPA via Docker. The latest stable image tag is openpolicyagent/opa:latest

You need to enter data and policy that can be later retrieved by the application.

Let us assume that "alice" is a paid customer and can create custom tiny url
Let us assume "bob" is a free user and he should not be able to create custom URL

samdeep$ cat myapi-acl.json 
{
 “alice”: [
  “read”,
  “write”
 ],
 “bob”: [
  “read”
 ]
}

samdeep$
samdeep$ cat myapi-policy.rego 
package myapi.policy
import data.myapi.acl
import input
default allow = false
allow {
    access = acl[input.user]
    access[_] == input.access
}
whocan[user] {
    access = acl[user]
    access[_] == input.access
}

Upload the data and policy to OPA Server running

samdeep$ curl -X PUT http://localhost:8181/v1/data/myapi/acl --data-binary @myapi-acl.json
samdeep$ curl -X PUT http://localhost:8181/v1/policies/myapi --data-binary @myapi-policy.rego

4. Make sure you have sqlite

Running the Application

1. Redis should be running
2. opa server should be running
2. Need SQlite for DB
3. The html is only for the anonymous user
4. REST API can be used to register a user, login, create URL etc

Here are the steps to run

1. Register an User (you can skip the step if you authenticate with google)
2. Login a User (Returns the token (Use Google Oauth2)
3. Use the access token for create url, redirect url, logout etc
4. Use the refresh token for refreshuser
