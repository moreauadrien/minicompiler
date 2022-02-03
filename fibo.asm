MOV R1, #temp_1
LDRB R0, [R1]
MOV R1, #var_i
STRB R0, [R1]

MOV R1, #temp_2
LDRB R0, [R1]
MOV R1, #var_max
STRB R0, [R1]

MOV R1, #temp_3
LDRB R0, [R1]
MOV R1, #var_un
STRB R0, [R1]

MOV R1, #temp_4
LDRB R0, [R1]
MOV R1, #var_un1
STRB R0, [R1]

MOV R1, #temp_5
LDRB R0, [R1]
MOV R1, #var_temp
STRB R0, [R1]

startwhile1
MOV R1, #var_i
LDRB R0, [R1]
MOV R1, #var_max
LDRB R3, [R1]
CMP R0, R3
MOV R0, #0x1
BCC condtrue2
MOV R0, #0x0
condtrue2
MOV R1, #temp_6
STRB R0, [R1]

MOV R1, #temp_6
LDRB R0, [R1]
MOV R3, #0x1
CMP R0, R3
BNE endwhile1
MOV R1, #var_un1
LDRB R0, [R1]
MOV R1, #var_temp
STRB R0, [R1]

MOV R1, #var_un
LDRB R0, [R1]
MOV R1, #var_un1
LDRB R3, [R1]
ADD R0, R0, R3
MOV R1, #temp_7
STRB R0, [R1]

MOV R1, #temp_7
LDRB R0, [R1]
MOV R1, #var_un1
STRB R0, [R1]

MOV R1, #var_temp
LDRB R0, [R1]
MOV R1, #var_un
STRB R0, [R1]

MOV R1, #var_i
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

B startwhile1
endwhile1
endprog END

temp_1 DCB 0x0
temp_2 DCB 0xA
temp_3 DCB 0x0
temp_4 DCB 0x1
temp_5 DCB 0x0
temp_6 DCB 0x0
temp_7 DCB 0x0
temp_8 DCB 0x1
temp_9 DCB 0x0
var_i DCB 0x0
var_max DCB 0x0
var_un DCB 0x0
var_un1 DCB 0x0
var_temp DCB 0x0
