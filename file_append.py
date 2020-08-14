file_name = '..\\Set\\SubtTV_2017_01_03_pcm.list.punct'
# file_name = '..\\Set\\SubtTV_2017_not_found'
extension = '.trn'

max_num = 80
file_list = []

for i in range(max_num):
  f = open(file_name + '_' + str((i+1)) + extension, 'rt', encoding = 'utf-8')
  file_list.append(f.readlines())

output = open(file_name + extension, 'wt', encoding = 'utf-8')
for i in range(max_num):
  # print(file_list[i])
  output.writelines(file_list[i])