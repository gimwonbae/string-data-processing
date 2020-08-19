import os
import re

check_f = open(os.path.join('..','Set','SubtTV_2017_01_03_pcm.list.punct.trn'), 'rt', encoding='utf-8')
org_f = open(os.path.join('..','Set','SubtTV_2017_01_03_pcm.list.trn'), 'rt', encoding='utf-8')

miss = open('miss.txt', 'wt', encoding ='utf-8')
match = open('match.txt', 'wt', encoding = 'utf-8')

check_list = check_f.readlines()
org_list = org_f.readlines()
re_exp = '[^가-힣0-9a-zA-Z]'

for check_line in check_list:
  check_line_split = check_line.split(' :: ')
  try:
    check_file = check_line_split[0]
    check_text = check_line_split[1]
  except:
    print(check_line)
    continue
  check_sub = re.sub(re_exp, '', check_text)

  i = 0
  for org_line in org_list:
    org_line_split = org_line.split(' :: ')
    org_file = org_line_split[0]
    
    if check_file == org_file:
      org_text = org_line_split[1]
      org_sub = re.sub(re_exp, '', org_text)
      if check_sub.lower() != org_sub.lower():
        miss.write(check_line)
      else:
        match.write(check_line)
      org_list = org_list[i+1:]
      break
    else:
      i+=1
  