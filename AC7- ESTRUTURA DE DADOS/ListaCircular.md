
# **Listas Circulares**

Uma lista circular é uma estrutura de dados em que cada nó aponta para o próximo nó na lista, e o último nó aponta de volta para o primeiro nó, formando um círculo.

## **Estrutura de um Nó**

Cada nó em uma lista circular simplesmente encadeada contém dois campos:

1. **Dados**: O valor ou informação armazenada no nó.
2. **Próximo**: Um ponteiro que aponta para o próximo nó na lista.

## **Estrutura da Lista**

A estrutura da lista contém apenas um campo:

1. **Ponteiro**: Um ponteiro do primeiro elemento da lista


### **1. Exibição dos Nós em uma Lista Circular**


- Começar com o primeiro nó 
- Imprimir o valor do nó atual.
- Ir para o próximo nó.
- Repitir os passos 2 e 3 até que o nó atual seja novamente o nó cabeça.
    
    ```
    Função ExibirListaCircular(nó):
        Se nó é NULL então
            Escrever "A lista está vazia"
            Retornar
    
        nóAtual <- nó
        Enquanto nóAtual.próximo ≠ nó
            Escrever nóAtual.dados
            nóAtual <- nóAtual.próximo
           
    
    Fim Função
    ```
    

### **2. Inserção de um Nó no Início da Lista**

A inserção no início de uma lista circular envolve a criação de um novo nó e o ajuste dos ponteiros para incluir esse nó na lista:

-Apontar o proximo nó do nó final para o novo nó

-Apontar o proximo nó do nó atual para o primeiro nó
    
   ```
nóFinal.proximoNó<-novoNó
novoNó.proximoNó<-nóInicial
novoNó<-novoNó.valor
        
    Fim Função
   ```
    

### **3. Exclusão do Primeiro Nó da Lista**

A exclusão do primeiro nó em uma lista circular requer ajustes nos ponteiros para remover o nó da lista:

- Se a lista tem apenas um nó (a cabeça que aponta para si mesma), simplesmente definir a cabeça como NULL.
- Encontrar o último nó da lista (o nó cujo próximo aponta para a cabeça).
- Fazer o próximo do último nó apontar para o segundo nó da lista (o próximo da cabeça atual).
- Atualizar a cabeça da lista para ser o segundo nó.
    
    ```
        Se nó.próximo é nó então
            // A lista tem apenas um nó
            nó.próximo = NULL
            Retornar NULL
        Senão
            nóAtual <- nó
            Enquanto nóAtual.próximo ≠ nó:
                nóAtual <- nóAtual.próximo
            Fim Enquanto  // No fim, o nóAtual é o último nó da lista circular (ele aponta para o nó inicial)
            
            nóAtual.próximo <- nó.próximo  // O último nó agora aponta para o segundo nó
            Liberar nó
            Retornar nóAtual.próximo  // O segundo nó agora é o primeiro
    Fim Função
    ```
