# API Gateway

## DEFINIÇÃO

É um modulo que tem como objetivo centralizar a entrada da sua rede e distribuir as requisições pro respectivos serviços internos da rede. Essa abordagem permite diversos beneficios mas também gera alguns pontos importantes que devem ser implementados para tomar o devido cuidado.

## VANTAGENS
Um unico ponto de entrada para as aplicações, dessa forma o cliente não precisa gravar o endereço de todas as suas aplicações, ele apenas vai gravar as de uma mesma url e o API Gateway irá encaminha para o serviço correto.

Ex: 
www.gateway-url.com/pagamento -> envia para o microserviço de pagamento.
www.gateway-url.com/news -> envia para o serviço de feed de noticias.

## DESVANTAGENS

Single point of failure (Unico ponto de falha): Utilizado para identificar 