
# Lista Duplamente Encadeada Ordenada

## Definição de Nó

Estrutura Lista:
 
 - Atual

Estrutura Nó:

 - Valor
 - PonteiroAnterior
 -  PonteiroProximo

## Algoritmo de Exibição

- Começar com o primeiro nó
- Imprimir o valor do nó 
- Ir para o próximo nó.
- Repetir os passos 2 e 3 até que o nó atual seja novamente o primeiro nó 
    
    ```
        Escrever nóAtual
        nóAtual<-PonteiroProximo
        Fazer até nóAtual=Null
         Escrever nóAtual
        nóAtual<-PonteiroProximo
    
    Fim Função

## Algoritmo de Busca
- Começar com o primeiro nó
- Procurar o Valor buscado no nóAtual
- Se não for o mesmo valor, ir para o próximo nó.
- Se for o mesmo valor , imprimir o seu nóAtual e Valor


```
Fazer até nóAtual=busca
	nóAtual<-proximoNó
	   se nóAtual=busca 
	   Escrever nóAtual e Valor
Fim 
```
## Algoritmo de Inserção

Procurar o ultimo nó
Procurar um espaço de memoria vazio
Apontar o proximoNó do ultimo nó para o espaço vazio
Inserir o valor no espaço de mémoria vazia
apontar o nóAnterior para o ultimo nó
```
Fazer
	nóAtual<-proximoNó 
		se proximoNó=null
			proximoNó<-memóriaVazia
			proximoNó.valor<-insert
			proximoNó.nóAnterior<-nóAtual
			Fim 
```
## Algoritmo de Remoção

Usar o algoritmo de busca para localizar o dado buscado
Apontar o nóAnterior.proximoNó<-nóAtual.proximoNó
Apontar o proximoNó.nóAnterior<-nóAtual.nóAnterior
Apagar o Valor do nóAtual
```
Fazer até nóAtual=busca
	nóAtual<-proximoNó
	   se nóAtual=busca 
	   nóAnterior.proximoNó<-nóAtual.proximoNó
	   proximoNó.nóAnterior<-nóAtual.nóAnterior
	   nóAtual.valor<-null
	   Fim
