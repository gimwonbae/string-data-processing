import re
# import time
import datetime

def to_list(file):
  f = open(file, 'rt', encoding = 'UTF-8')
  lines = f.readlines()
  return lines

w_punct = to_list('broadcast_w_punct_u_double.txt')
backup = w_punct

f = open('SubtTV_2017_01_03_pcm.list.trn', 'rt', encoding = 'UTF-8')
new_f = open('SubtTV_2017_01_03_pcm.list.punct.trn', 'wt', encoding = 'UTF-8')
not_found = open('SubtTV_2017_not_found', 'wt', encoding = 'UTF-8')

print(datetime.datetime.now())
while True:
  # counter = time.time()

  line = f.readline()
  if not line : break

  file_name = line.split(' :: ')[0]
  org_txt = line.split(' :: ')[1][:-1]
  
  flag = True
  org_txt_sub = re.sub('[ ]', '', org_txt)    
  first_word = org_txt.split(' ')[0]
  second_word = org_txt.split(' ')[-1]
  # print('org : ', org_line_sub)

  for cmp_line in w_punct:
    cmp_txt = cmp_line[:-1]
    cmp_txt_sub = re.sub('[,.?!~ ]', '', cmp_txt)
    # print('cmp : ', cmp_line_sub)

    if org_txt_sub in cmp_txt_sub:
      start = cmp_txt.find(first_word)
      check = cmp_txt.find(second_word, start + len(first_word))
      end = cmp_txt.find(' ', check)
      if end == -1:
        new_f.write(file_name + ' :: ' + cmp_txt[start:] + '\n')
      else:
        new_f.write(file_name + ' :: ' + cmp_txt[start:end] + '\n')
      flag = False
      # print('done : ', time.time() - counter)
      w_punct = w_punct[w_punct.index(cmp_line)+1:]
      break
  if flag:
    not_found.write(line)
    w_punct = backup
    # print('fail : ', time.time() - counter)