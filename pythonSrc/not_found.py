import os
import re

def view_list(line, path, type):
  # f = open(line, rt, encoding= encodings)
  org_path = line.split(' :: ')[0]
  file_name = org_path.split('/')[-2]
  if type == 'text':
    f = open(os.path.join(path,file_name,'data',file_name+'.text'), 'rt', encoding = 'utf-8')
  elif type == 'tok':
    f = open(os.path.join(path,file_name,'data',file_name+'.tok'), 'rt', encoding = 'cp949')
  else:
    print('type = text or tok')
  return f.readlines()

# print(view_list('SubtTV_Database/2017/03/KBS1/KBS1_2017_0301_0026/KBS1_2017_0301_0026_243_003.pcm :: 규모 오점팔의 경주 지진이 난 지 오개월이 넘었지만 여진은 지금도 진행 중입니다', "C:\\Project\\20200804.-방송DB후처리\\broadcast_text\\KOR", 'tok'))

f = open(os.path.join('..', 'Set', 'SubtTV_2017_01_03_not_found.trn'), 'rt', encoding = 'utf-8')
re_exp = '[^a-zA-Z가-힣0-9]'
dir_path = '..\\broadcast_text\\KOR'

# i = 0
for line in f.readlines():
  text_list = view_list(line, dir_path, 'text')
  tok_list = view_list(line, dir_path, 'tok')

  line = line.split(' :: ')[1]
  line_sub = re.sub(re_exp, '', line).lower()
  # print(line_sub)
  index = 0
  for tok in tok_list:
    tok_sub = re.sub(re_exp, '', tok).lower()
    # print(tok_sub)
    if line_sub in tok_sub :
      break
    index +=1
  try:
    print(line, 'text : ', text_list[index], 'tok  : ',tok_list[index].replace('] ','').replace(' +','').replace('+','').replace('[',''))
  except:
    print('!!! ', line)