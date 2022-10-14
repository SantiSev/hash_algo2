package diccionario

import (
	"diccionario/lista"
	"fmt"
)

type hashDato[K comparable, V any] struct {
	clave K
	valor V
}

type hashMap[K comparable, V any] struct {
	hashArray []lista.Lista[hashDato[K, V]] //todo O(1) tonces todo bien 8)
	longitud  int
}

func (h hashMap[K, V]) convertir(T any) int {
	dato := convertirABytes[K](T)
	index := h.sdbmHash(dato)
	return index
}

func convertirABytes[K comparable](clave any) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h hashMap[K, V]) sdbmHash(data []byte) int {
	// documentacion: https://www.programmingalgorithms.com/algorithm/sdbm-hash/c/
	var hash uint64

	for _, b := range data {
		hash = uint64(b) + (hash << 6) + (hash << 16) - hash
	}

	return int(hash) % len(h.hashArray)
}

func (h *hashMap[K, V]) Guardar(clave K, valor V) {
	nuevoDato := &hashDato[K, V]{clave: clave, valor: valor}
	index := h.convertir(clave)
	h.hashArray[index].InsertarPrimero(*nuevoDato)
	h.longitud++
}

func (h hashMap[K, V]) Pertenece(clave K) bool {
	index := h.convertir(clave)
	listaIndex := h.hashArray[index]
	if listaIndex.EstaVacia() {
		return false
	} else {
		for iter := listaIndex.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			if iter.VerActual().clave == clave {
				return true
			}
		}
	}
	return false
}

func (h hashMap[K, V]) Obtener(clave K) V {
	index := h.convertir(clave)
	subLista := h.hashArray[index]
	if subLista.EstaVacia() {
		panic("La clave no pertenece al diccionario")
	}
	for iter := subLista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().clave == clave {
			return iter.VerActual().valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hashMap[K, V]) Borrar(clave K) V {
	index := h.convertir(clave)
	subLista := h.hashArray[index]
	if subLista.EstaVacia() {
		panic("La clave no pertenece al diccionario")
	}
	for iter := subLista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().clave == clave {
			dato := iter.Borrar()
			h.longitud--
			return dato.valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (h hashMap[K, V]) Cantidad() int {
	return h.longitud
}

func (h hashMap[K, V]) Iterar(f func(clave K, valor V) bool) {

	for _, subLista := range h.hashArray {
		if !subLista.EstaVacia() {
			for iter := subLista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
				dato := iter.VerActual()
				if !f(dato.clave, dato.valor) {
					return
				}
			}
		}
	}
}

func (h hashMap[K, V]) Iterador() IterDiccionario[K, V] {
	//TODO implement me
	panic("implement me")
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	h := new(hashMap[K, V])
	h.hashArray = make([]lista.Lista[hashDato[K, V]], 90)
	for i := range h.hashArray {
		h.hashArray[i] = lista.CrearListaEnlazada[hashDato[K, V]]()
	}
	return h
}

/*
		    .  ."|
   /| /  |  _.----._
  . |/  |.-"        ".  /|
 /                    \/ |__
|           _.-"""/        /
|       _.-"     /."|     /
 ".__.-"         "  |     \
    |              |       |
    /_      _.._   | ___  /
  ."  ""-.-"    ". |/.-.\/
 |    0  |    0  |     / |
 \      /\_     _/    "_/
  "._ _/   "---"       |
  /"""                 |
  \__.--                |_
    )          .        | ".
   /        _.-"\        |  ".
  /     _.-"             |    ".
 (_ _.-|                  |     |"-._.
   "    "--.             .J     _.-'
           /\        _.-" | _.-'
          /  \__..--"   _.-'
         /   |      _.-'
        /| /\|  _.-'
       / |/ _.-'
      /|_.-'
    _.-'
*/
