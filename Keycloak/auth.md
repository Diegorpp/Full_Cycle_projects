Autenticação e Autorização

Resource Owner - Dono do recurso

Client - Cliente que quer usar o recurso.

Resource Server - Servidor que vai perguntar se aquele cliente pode ou não usar fazer algo

Authorization Server - Responsável pro definir se aquele cliente possui de fato permissão para realizar.

Oauth 2 - Framework de autorização.
AUTORIZAÇÃO E NÃÃÃO AUTENTICAÇÃO.

![oauth](images/oauth.png)

OpenID Connect - Complemento para o Oauth, usado como uam camada que funciona em cima dele.

Utiliza um autentication endpoint no qual ao passar um response_type ele vai alinhar outras execuções. Segue alguns parâmetros esperados nesse campo:
code, token, id_token, id_token token, code id_token, code token id_token e none

id_token: Possui as informações necessarias para garantir a autenticação do usuário, como ID, Email e iformações adicionais passadas pelo servidor de autenticação.
- JWT (Json Web Token)

    - Header
    - Payload
    - Signature

![jwt](images/jwt.png)


