# jwt-user

Good References

https://github.com/victorsteven/jwt-best-practices

https://github.com/ishrivatsa/user

https://github.com/open-policy-agent/opa

https://github.com/open-policy-agent/example-api-authz-go

https://livebook.manning.com/book/microservices-security-in-action/welcome/v-7/

https://medium.com/@cgrant/golang-oauth2-google-example-for-web-and-api-59187ce8b119

Requirements
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
