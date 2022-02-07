MOV R1, #temp_1
LDRB R0, [R1]
MOV R1, #var_x
STRB R0, [R1]

MOV R1, #temp_2
LDRB R0, [R1]
MOV R1, #var_y
STRB R0, [R1]

startwhile1
MOV R1, #temp_3
LDRB R0, [R1]
MOV R1, #var_x
LDRB R3, [R1]
CMP R0, R3
MOV R0, #0x1
BCC condtrue2
MOV R0, #0x0
condtrue2
MOV R1, #temp_4
STRB R0, [R1]

MOV R1, #temp_4
LDRB R0, [R1]
MOV R3, #0x1
CMP R0, R3
BNE endwhile1
MOV R1, #var_y
LDRB R0, [R1]
MOV R1, #var_y
LDRB R3, [R1]
ADD R0, R0, R3
MOV R1, #temp_5
STRB R0, [R1]

MOV R1, #temp_5
LDRB R0, [R1]
MOV R1, #var_y
STRB R0, [R1]

MOV R1, #var_x
LDRB R0, [R1]
MOV R1, #temp_6
LDRB R3, [R1]
SUB R0, R0, R3
MOV R1, #temp_7
STRB R0, [R1]

MOV R1, #temp_7
LDRB R0, [R1]
MOV R1, #var_x
STRB R0, [R1]

B startwhile1
endwhile1
endprog 
END
B endprog
temp_1 DCB 0x9
temp_2 DCB 0xA
temp_3 DCB 0x0
temp_4 DCB 0x0
temp_5 DCB 0x0
temp_6 DCB 0x1
temp_7 DCB 0x0
var_x DCB 0x0
var_y DCB 0x0
