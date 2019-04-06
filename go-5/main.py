import hashlib
import json
import string
import string
import requests

alfabeto = list(string.ascii_lowercase)

def get_token():
    return "48c62d18afa2e20201ff6c7f6130857ba7f6a5f1"


def save_file(arquivo, data):
    with open(arquivo, "w") as outfile:
        json.dump(data, outfile, indent=4)


def get():
    r = requests.get(
        f"https://api.codenation.dev/v1/challenge/dev-ps/generate-data?token={get_token()}"
    )
    return r.json()


def espaceLetra(codigo):
    d = [32, 46, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57]
    if codigo in d:
        return True
    else:
        return False


def get_sha1(texto):
    texto = texto.encode("utf-8")
    return hashlib.sha1(texto).hexdigest()


def criptografa(texto, deslocamento):
    elementos = list(texto.lower())
    for k, v in enumerate(elementos):
        if espaceLetra(ord(v)):
            pass
        else:
            letra = ord(v)
            elementos[k] = chr(letra + deslocamento)
    return "".join(elementos)

def criptografa2(texto, deslocamento):
    elementos = list(texto.lower())
    cifrado = ''
    
    for v in elementos:
        if espaceLetra(ord(v)):
            cifrado += v
        else:
             indice  =  alfabeto.index(v) 
             if (indice + deslocamento) % 26 > 0 and (indice + deslocamento) % 26 < 25:
                c=(indice + deslocamento) % 26
                cifrado += alfabeto[c]
                        


def descriptografa(texto, deslocamento):
    elementos = list(texto.lower())
    for k, v in enumerate(elementos):
        if espaceLetra(ord(v)):
            pass
        else:
            letra = ord(v)
            elementos[k] = chr((letra - deslocamento))
    return "".join(elementos)

def descriptografa2(texto, deslocamento):
    elementos = list(texto.lower())
    decifrado = ''
    
    for v in elementos:
        if espaceLetra(ord(v)):
            decifrado += v
        else:
             indice  =  alfabeto.index(v) 
             if (indice + deslocamento) % 26 > 0 and (indice + deslocamento) % 26 < 25:
                c=(indice - deslocamento) % 26
                decifrado += alfabeto[c]
                        
    return decifrado



if __name__ == "__main__":
    r = get()
    save_file("answer.json", r)
    d = descriptografa2(r["cifrado"], r["numero_casas"])
    r["decifrado"] = d
    save_file("answer.json", r)
    r["resumo_criptografico"] = get_sha1(d)
    save_file("answer.json", r)
