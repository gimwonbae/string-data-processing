f = open('broadcast_w_punct.txt', 'rt', encoding = 'UTF-8')
new_f = open('broadcast_w_punct_u_double.txt', 'wt', encoding = 'UTF-8')
prev = ''
l = list()
while True:
  line = f.readline()
  if not line: break
  if prev != line:
    l.append(line)
    prev = line
new_f.writelines(l)