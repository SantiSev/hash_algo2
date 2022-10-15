package diccionario

import (
	"diccionario/lista"
	"fmt"
	"math"
)

const (
	LONGITUD_INICIAL     = 89
	REDIMENSION_AGRANDAR = 10
	MINIMO_REDIMENSION   = 4
)

type hashDato[K comparable, V any] struct {
	clave K
	valor V
}

type hashMap[K comparable, V any] struct {
	hashArray []lista.Lista[hashDato[K, V]] //todo O(1) tonces todo bien 8)
	longitud  int
}

type iteradorHash[K comparable, V any] struct {
	hashEstructura []lista.Lista[hashDato[K, V]]
	index          int
	subListaIter   lista.IteradorLista[hashDato[K, V]]
}

// Implementacion de HashMap

// TODO sacar la func de hashing (y compania) del tda.
func (h hashMap[K, V]) convertir(T any) int {
	dato := convertirABytes[K](T)
	index := h.sdbmHash(dato)
	if index < 0 {
		index *= -1
	}
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

func (h *hashMap[K, V]) actualizar(clave K, valorActualizado V) {
	index := h.convertir(clave)
	listaIndex := h.hashArray[index]
	for iter := listaIndex.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().clave == clave {
			iter.Borrar()
			iter.Insertar(hashDato[K, V]{clave: clave, valor: valorActualizado})

		}
	}
}

func (h *hashMap[K, V]) redimensionar(valorARedimensionar int) {
	redim := new(hashMap[K, V])
	redim.hashArray = make([]lista.Lista[hashDato[K, V]], valorARedimensionar)
	for i := range redim.hashArray {
		redim.hashArray[i] = lista.CrearListaEnlazada[hashDato[K, V]]()
	}
	redim.longitud = h.longitud
	for _, subLista := range h.hashArray {
		for iter := subLista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			datoActual := iter.VerActual()
			redim.Guardar(datoActual.clave, datoActual.valor)
		}
	}
	h = redim
}

func (h *hashMap[K, V]) Guardar(clave K, valor V) {
	nuevoDato := &hashDato[K, V]{clave: clave, valor: valor}
	index := h.convertir(clave)
	if h.Pertenece(clave) {
		h.actualizar(clave, valor)
		return
	}

	h.hashArray[index].InsertarPrimero(*nuevoDato)

	if h.hashArray[index].Largo() >= REDIMENSION_AGRANDAR {
		h.redimensionar(proxPrimo(h.longitud * 2))
	}
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

func (h *hashMap[K, V]) Obtener(clave K) V {
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

func (h hashMap[K, V]) Borrar(clave K) V {
	index := h.convertir(clave)
	subLista := h.hashArray[index]
	if subLista.EstaVacia() {
		panic("La clave no pertenece al diccionario")
	}
	for iter := subLista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		if iter.VerActual().clave == clave {
			dato := iter.Borrar()
			if MINIMO_REDIMENSION*h.Cantidad() <= len(h.hashArray) && MINIMO_REDIMENSION*h.Cantidad() >= LONGITUD_INICIAL {
				h.redimensionar(proxPrimo(h.longitud / 2))
			}
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
	iter := new(iteradorHash[K, V])
	iter.hashEstructura = h.hashArray
	return iter
}

// Implementacion de iter Externo

func (i iteradorHash[K, V]) HaySiguiente() bool {
	//TODO implement me
	panic("implement me")
}

func (i iteradorHash[K, V]) VerActual() (K, V) {
	//TODO implement me
	panic("implement me")
}

func (i iteradorHash[K, V]) Siguiente() K {
	//TODO implement me
	panic("implement me")
}

// CrearHash + Otras funciones privadas

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	h := new(hashMap[K, V])
	h.hashArray = make([]lista.Lista[hashDato[K, V]], 89)
	for i := range h.hashArray {
		h.hashArray[i] = lista.CrearListaEnlazada[hashDato[K, V]]()
	}
	return h
}

func esPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i < int(math.Sqrt(float64(n))+1) {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func proxPrimo(n int) int {
	if n <= 1 {
		return 2
	}
	primo := n
	encontrado := false
	for !encontrado {
		primo += 1
		if esPrimo(primo) {
			encontrado = true
		}
	}
	return primo
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
