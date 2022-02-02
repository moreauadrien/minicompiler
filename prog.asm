MOV R1, #temp_1
LDRB R0, [R1]
MOV R1, #var_i
STRB R0, [R1]

MOV R1, #var_i
LDRB R0, [R1]
MOV R1, #temp_2
LDRB R3, [R1]
ADD R0, R0, R3
MOV R1, #temp_3
STRB R0, [R1]

MOV R1, #temp_3
LDRB R0, [R1]
MOV R1, #temp_4
LDRB R3, [R1]
ADD R0, R0, R3
MOV R1, #temp_5
STRB R0, [R1]

MOV R1, #temp_5
LDRB R0, [R1]
MOV R1, #temp_6
LDRB R3, [R1]
ADD R0, R0, R3
MOV R1, #temp_7
STRB R0, [R1]

MOV R1, #temp_7
LDRB R0, [R1]
MOV R1, #temp_8
LDRB R3, [R1]
ADD R0, R0, R3
MOV R1, #temp_9
STRB R0, [R1]

MOV R1, #temp_9
LDRB R0, [R1]
MOV R1, #var_i
STRB R0, [R1]

MOV R1, #var_i
LDRB R0, [R1]
MOV R1, #temp_10
LDRB R3, [R1]
SUB R0, R0, R3
MOV R1, #temp_11
STRB R0, [R1]

MOV R1, #temp_11
LDRB R0, [R1]
MOV R1, #var_i
STRB R0, [R1]

endprog END

temp_1 DCB 0x6
temp_2 DCB 0x1
temp_3 DCB 0x0
temp_4 DCB 0x5
temp_5 DCB 0x0
temp_6 DCB 0xE
temp_7 DCB 0x0
temp_8 DCB 0x12
temp_9 DCB 0x0
temp_10 DCB 0x9
temp_11 DCB 0x0
var_i DCB 0x0
