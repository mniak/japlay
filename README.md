JAPlayer
=========

Player multiplataforma para as músicas no banco de dados do LouvorJA.

**Ainda em fase de protótipo**

## Como rodar no Linux

Como este é um protótipo, quase nada foi implementado ainda, inclusive o mecanismo de download e de gerenciamento de coletâneas e arquivos.

### Você vai precisar de:
- Conhecimento básico de terminal do Linux
- Cópia da pasta `config` do aplicativo Louvor JA normal. Esta pasta pode ser encontrada no windows em `C:\Program Files (x86)\LouvorJA\config`.
- Uma versão adaptada do banco de dados: https://raw.githubusercontent.com/mniak/japlay/master/assets/DB.db
- Executável `japlay`, que é o programa que vamos testar. https://github.com/mniak/japlay/releases/download/v0.0.1/japlay-v0.0.1-linux.tar.gz
    - Este é um arquivo compactado, como se fosse um .zip. Para descompactar o arquivo, abra um terminal no linux, navegue até a pasta e rode o comando `tar -xvf japlay-v0.0.1-linux.tar.gz`.
- Bibliotecas **SDL2**, **SDL2_ttf**, **SDL2_image** e **SDL2_mixer** instaladas no se linux. No ubuntu podem ser instaladas facilmente rodando o comando `sudo apt-get install libsdl2{,-image,-mixer,-ttf}`


### Preparando o ambiente
Crie uma nova pasta, por exemplo chamada `teste-ja-play` e copie a pasta `config` para dentro dela.
Também copie para dentro dela o arquivo `japlay`.

Em seguida, abra um terminal, navegue até a pasta `teste-ja-play` e mande tocar o hino nº 1 do hinário 2022 com o seguinte comando:
```
./japlay play 1
```


