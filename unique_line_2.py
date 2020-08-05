f = open('broadcast_w_punct.txt', 'rt', encoding = 'UTF-8')
new_f = open('broadcast_w_punct_u.txt', 'wt', encoding = 'UTF-8')
l = f.readlines()
input_dic = {}
r_list = []

for i, v in enumerate(l):
  get_value = input_dic.get(v, None)
  if get_value == None:
    input_dic[v] = i
    r_list.append(v)

new_f.writelines(r_list)