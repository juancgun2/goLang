# Seminario GoLang

## Description 
Este proyecto analiza cadenas de strings y de ser correctas,
devuelve una estructura indicando tipo, longitud y valor de la cadena. 

## Implementation 
Definimos una estructura con tres atributos y una serie de comportamientos, 
que hacen al analisis de la cadena y la creacion de la estructura. 
Utilizamos regexp para buscar numeros en un string y un parse (string to int) para buscar letras en un string que supone ser numerico.  

## Input
Se espera una cadena con el siguiente formato: 
    { TX04JUAN } 
Donde los primeros 2 caracteres son el tipo,
los siguientes 2 indican longitud del valor
y por ultimo el valor.

Por el momento el sistema soporta los siguientes tipos: 
- TX = TEXTO 
- NN = NUMERICO



