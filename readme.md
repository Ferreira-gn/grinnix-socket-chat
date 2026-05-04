# Objetivo :
Realizar a criação de um chat de mensagens orquestrado por um servidor socket, onde os usuários conectados possam se comunicar através de mensagens globais e particulares.

## Requisitos a serem compridos :

- Como usuário quero me conectar em um chat 
- Como usuário quero poder visaulizar as mensagens que estão ocorrendo no chat
- Como usuáio quero poder enviar mensagens no chat 
- Como usuário quero poder enviar mensagens para todos ( @everony )
- Como usuário quero poder "susurrar" uma mensagem para alguém ( @name )

## Regras de negócio :

- Mensagens sem @everony já são por padrão globais
- Mensagens com o @everony serão mandadas para todos os usuário mas com destaque maior
- Mensagens susurradas só podem ser vistas por que as envio e quem as recebeu