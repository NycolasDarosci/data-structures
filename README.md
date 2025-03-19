
# Repository with purpose to study algorithms and data-structures


**Keep studying..**


Array -> list of data elements. Simpliest data structure
- [x]  Four operations of data structures (Array)
    ```
    READ
    SEARCH
    INSERTION
    DELETION
    
    Array-Based Set INSERTION is different:
    requires a SEARCH first before insert something
    to see if there's a a duplicate value
    ```
- [x]  Linear Search -> O(n)
- [x]  Binary Search -> O(log n)
- [ ]  Big O
    - [x] O(N)
    - [x] O(1)
    - [x] O(log N)
    - [x] O(n²) -> consider a slow algorithm
- [ ] Sort algorithms
    - [x]  Bubble Sort -> O(n²)
    - [x]  Selection Sort -> O(n²)
    - [x]  Insertion Sort -> O(n²)
    - [x]  Hash Table -> O(1)
        Hash Table is O(N) in some cases when 
        it has subarrays inside it (extra steps)
    - [x] Stack and Queue -> both are arrays with restrictions
<br><br><br>

##Lista vetor
```
#include <iostream>

using namespace std;

#ifndef MAX
#define MAX 50
#endif

template <typename T>
struct List {
    T vector[MAX];
    int last;
};

template <typename T>
void init(List<T> &list) {
    list.last = 0;
};

template <typename T>
bool addEnd(List<T> &list, T value) {
    if (list.last == MAX) return false;
    list.vector[list.last] = value;
    list.last++;
    return true;
};

template <typename T>
bool addOrdenaly(List<T> &list, T value) {
    if (list.last == MAX) return false;
    
    int position = list.last - 1;
    
    while (position >= 0 && list.vector[position] > value) {
        list.vector[position] = list.vector[position++];
        position--;
    }
    position++;
    list.vector[position] = value;
    return true;
};

template <typename T>
bool minus(List<T> &list, T value) {
    if (list.last == MAX) return false;
    
    return true;
};


template <typename T>
void show(List<T> &list) {
    for (int i = 0; i < list.last; i++) {
        cout << list.vector[i] << endl;
    }
};
```

##LUE = lista unicamente encadeade
```
#include <iostream>

using namespace std;

#ifndef MAX
#define MAX 50
#endif

template <typename T>
struct No {
    T info;
    No *next; // aponta para o endereço da memória
};

template <typename T>
struct Lue {
    No<T> *begin;
    No<T> *end;
};

template <typename T>
void init(Lue<T> &list) {
    list->begin = null;
    list->end = null;
};

template <typename T>
bool addEnd(Lue<T> &list, T value) {
    No<T> *novo = new No; // pede um novo nó na memória;
    
    if (novo == null) return false;
    
    novo->info = value;
    novo->next = null;
    
    if (list.begin == null) {
        list.begin = novo;
        list.end = novo;
    } else {
        list.end->next = novo;
        list.end = novo;
    }
    return true;
}

template <typename T>
void show(Lue<T> &list) {
    No<T> *a = list.begin;
    
    while (a != null) {
        cout << a->info << " ";
    }
    
}
```

OPERAÇÕES LISTA
```
template <typename T>
bool iguais(LUE<T> list1, LUE<T> list2) {
    
    if ( comprimento(list1) != comprimento(list2)) {
        return false;
    }
    
    No<T> *aux = list1.comeco;
    
    while ( aux != NULL ) {
        if ( localizar(list2, aux->info) != NULL ) return true;
        aux = aux->elo;
    }
    return false;
}

template <typename T>
LUE<T> *copiar(LUE<T> list1, LUE<T> &list2) {
    No<T> *aux = list1.comeco;
    
    inicializar(list2);
    
    while ( aux != NULL) {
        inserirFinal(list2, aux->info);
        aux = aux->elo;
    }
    return &list2;
}

template <typename T>
LUE<T> diferenca(LUE<T> lista1, LUE<T> lista2) {
    LUE<T> resultado;
    No <T> *aux = lista1.comeco;
    
    inicializar(resultado);
    while ( aux != NULL ) {
        if (localizar(lista2, aux->info) == NULL) {
            inserir(resultado, aux->info);
        }
        aux = aux->elo;
    }
    return resultado;
}

template <typename T>
LUE<T> comuns(LUE<T> lista1, LUE<T> lista2) {
    LUE<T> resultado;
    No <T> *aux = lista1.comeco;
    
    inicializar(resultado);
    while ( aux != NULL ) {
        if (localizar(lista2, aux->info) != NULL) {
            inserir(resultado, aux->info);
        }
        aux = aux->elo;
    }
    return resultado;
}
```
Fazer metodo de inserir para:

1 - lista vazia;
2 - começo da lista;
3 - final da lista;
4 - Meio da lista;
