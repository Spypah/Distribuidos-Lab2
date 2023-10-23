# Grupo01-Laboratorio-2

Levanto datanode
Levanto OMS
Levanto ONU
Levanto continentes

Conexiones
**Continentes** Se conectan a la ip de la OMS
**OMS** Se conecta a la ip de los Datanodes
**ONU** Se conecta a la ip de la OMS
**Datanodes** No se conecta


1 iteración 5 datos
luego 1 dato cada 3 segundos
**Continentes** Obtienen al azar un nombre del TXT y determinan estado de la persona.
**Continentes** Se conectan de manera sincrona a la OMS para ingresar información de las personas (nombres y estados)
**OMS** Obtiene la info
**OMS** Registra nueva linea dentro del archivo DATA.txt que contiene, un ID dado por el NameNode, el número del DataNode y el estado de la persona
**OMS** Envia un mensaje al DataNode con el ID asignado por el NameNode, el nombre y el apellido de la persona
**DataNode** Almacena la información entregada por NameNode y lo guarda en DATA.txt

fin iteracion

Interacciones
**ONU** Ejecuta una query por estado de personas a OMS
**OMS** Pregunta por las personas al datanode
**Datanode** Cuando recibe una query de la OMS por id, retorna nombre y apellido desde DATA.txt
**OMS** Retorna información a la ONU
**ONU** Muestra por pantalla el resultado


