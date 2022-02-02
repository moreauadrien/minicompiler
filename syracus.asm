test  
		
MOV R1, #vnn_i
LDRB R0, [R1]
MOV R3, #temp_1
CMP R0, R3
BEQ fin
		
corps 
MOV R3, #temp_1
SUB R0, R0, R3
STRB R0, [R1]
MOV R1, #res
LDRB R0, [R1]
MOV R3, #temp_1
AND R0, R0, R3
CMP R0, R3
BEQ impaire
BNE paire
		
paire
MOV R1, #res_i
LDRB R0, [R1]
MOV R3, #temp_1
LSR R0, R0, R3
STRB R0, [R1]
B test
		
impaire
MOV R1, #res_i
LDRB R0, [R1]
MOV R3, R0
ADD R0, R0, R0
ADD R0, R0, R3
ADD R0, R0,  #temp_1
STRB R0, [R1]
B test
		
		
fin
B fin
vnn_i DCB 0x06
res_i DCB 0x0E
temp_1 DCB 0x01