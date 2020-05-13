# Reglas

Tipos de frase
---

* VERBOS
   1. TIEMPOASPECTOMODO + JUEGO[A/B] + VERBO
* PRONOMBRE + SUSTANTIVO
* PREPOSICIÓN + SUSTANTIVO
* DEMOSTRATIVO + SUSTANTIVO
* VERBO + PRONOMBRE + SUSTANTIVO
* VERBO + COMPLEMENTO + SUSTANTIVO

**Estructura de Reglas**
---
Estructura para prioridad de reglas. Todos los TAG que devuelve linguakit poseen la propiedad Type, a excepción: 
* [I]=Interjections 
* [CARD]=Cardinales y
* [DATE]=fecha y hora
```
kaqchikel_rules_collection = [
    {
        "elements":"1",
        "pos0":{
            "tag":"VERB",
            "type":"M"
        }
    },
    {
        "elements":"2",
        "pos0":{
            "tag":"DET",
            "type":"P"            
        },
        "pos1":{
            "tag":"NOUN",
            "type":"C"
        }
    },
    {
        "elements":"2",
        "pos0":{
            "tag":"DET",
            "type":"D"
        }, 
        "pos1":{
            "tag":"NOUN",
            "type":"C"
        } 
    },
    {
        "elements":"2",
        "pos0":{
            "tag":"PRP",
            "type":"P"
        }
        "pos1":{
            "tag":"NOUN",
            "type":"C"
        }
    },
    {
        "elements":"3",
        "pos0:"{
            "tag":"DET",
            "type":"P"
        },
        "pos1:"{
            "tag":"NOUN",
            "type":"C"
        },
        "pos2:"{
            "tag":"VERB",
            "type":"M"
        }
    },
    {
        "elements":"5",
        "pos0":{
            "tag":"NOUN",
            "type":"C"
        },
        "pos1":{
            "tag":"NOUN",
            "type":"P"
        },
        "pos2":{
            "tag":"VERB",
            "type":"M"
        },
        "pos3":{
            "tag":"DET",
            "type":"I"
        },
        "pos4":{
            "tag":"NOUN",
            "type":"C"
        }
    }
    //HASTA N REGLAS
]
```

**Estructura Verbos**
---
Este análisis se realizará cuando Linguakit nos de el valor:
* **[VERB]**

Para construir la regla de verbos necesitamos almacenar los prefijos de:
* Tiempo, Aspecto y Modo
* Pronombres para verbos tipo transitivo e intransitivo 
* Información del verbo

Tiempo, Aspecto, Modo
```
kaqchikel_time_collection = {
    "pasado":"x",
    "pasado_def_1":"k",
    "pasado_def_2":"t"
    "presente_1":"y"
    "presente_2":"n",
    "futuro_1":"xk",
    "futuro_2":"xt",
    "futuro_inmediato_1":"y",
    "futuro_inmediato_2":"n"
}
```

Marcadores de Persona - Verbos Transitivos
```
kaqchikel_person_transitive_mark_collection = {
    "primera":{
        "singular":{
            "consonante_1":"nu",
            "consonante_2":"in-",
            "vocal_1":"w",
            "vocal_2":"inw"
        },
        "plural":{
            "consonante":"a",
            "vocal":"qa"
        }
    }, 
    "segunda":{
        "singular":{
            "consonante":"a",
            "vocal":"aw"
        },
        "plural":{
            "consonante":"i",
            "vocal":"iw"
        }
    },
    "tercera":{
        "singular":{
            "consonante_1":"u",
            "consonante_2":"ru",
            "vocal":"r"
        },
        "plural":{
            "consonante":"ki",
            "vocal":"k"
        }
    }
}
```
Marcadores de Persona - Verbos Intransitivos
```
kaqchikel_person_intransitive_mark_collection = {
    "primera":{
        "singular":{
            "consonante_1":"i",
            "consonante_2":"in",
            "vocal":"in"
        },
        "plural":{
            "consonante":"oj",
            "vocal":"oj"
        }
    },
    "segunda":{
        "singular":{
            "consonante_1":"a",
            "consonante_2":"at",
            "vocal":"at"
        },
        "plural":{
            "consonante":"ix",
            "vocal":"ix"
        }
    },
    "tercera":{
        "singular":{
            "consonante":"0/",
            "vocal":"0/"
        },
        "plural":{
            "consonante":"e",
            "vocal":"e'"
        }
    }
}
``` 
Verbo: su valor se obtiene de la estructura de traducciones
```
{
    "spanish":"word_spanish",
    "kaqchikel":"word_kaqchikel"
}
```

---
**Frases Varias**
---

**Pronombres y Sustantivos**

Este análisis se realiza cuando Linguakit nos de los resultados de:

* **[DET]**
  * Type:P (possesive)
```
  * Marcadores de Persona - Verbos Transitivos
```
* **[NOUN]**
```
Sustantivo: su valor se obtiene de la estructura de traducciones
```
