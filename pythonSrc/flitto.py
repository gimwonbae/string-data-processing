import pandas as pd
import os
import shutil

def allFile(path):
  res = []
  for root, dirs, files in os.walk(path):
    for file in files:
      res.append(os.path.join(root, file))
  return res

def allDir(path):
  res = []
  for root, dirs, files in os.walk(path):
    for Dir in dirs:
      res.append(os.path.join(root, Dir))
  return res

def selectFiles(files, condition):
  newfiles = []
  for file in files:
    if condition in file:
      newfiles.append(file)
      # print(file)
  return newfiles

def toDF(path):
  files = allFile(path)
  dfs= pd.DataFrame()
  for file in files:
    if (".new.xlsx" in file) and ("$" not in file):
      df = pd.read_excel(file).iloc[:,0:6]
      dfs = dfs.append(df)
  return dfs.dropna(axis=0)

def nameRule(name):
  return name.replace("(","").replace(")","").replace(" ","").replace("-","").replace("_","")

def makeFile(files, extension, i, blackList, row):
  flag = True
  change = False
  for file in files:
    if nameRule(row.파일명) == nameRule(os.path.basename(file).replace(extension, '')):
      flag = False
      if file not in blackList:
        change = True
        blackList.append(file)
        fileName = file.split(os.path.sep)[-2] + '_' + str(i).zfill(5) + extension
        fullName = os.path.join(os.path.dirname(file), fileName)
        newName = fullName.replace('20200705_Flitto_Rantacar_samples','Filtto_Rantacar')
        shutil.copy(file, newName)
        break
  if flag:
    print(row.파일명 + ' : ' + extension + 'not found')
  return [blackList, change]
  
def makeTxt(files, extension, i, blackList, row, en):
  flag = True
  for file in files:
    if nameRule(row.파일명) == nameRule(os.path.basename(file).replace(extension, '')):
      flag = False
      if file not in blackList:
        blackList.append(file)
        fileName = file.split(os.path.sep)[-2] + '_' + str(i).zfill(5)
        koName = os.path.join(os.path.dirname(file), fileName + '_ko.txt').replace('20200705_Flitto_Rantacar_samples','Filtto_Rantacar')
        enName = os.path.join(os.path.dirname(file), fileName + '_en.txt').replace('20200705_Flitto_Rantacar_samples','Filtto_Rantacar')
        shutil.copy(file, koName)

        f = open(enName, 'wt')
        f.write(en[str(row.파일명)])
        break
  if flag:
    print(row.파일명 + ' : ' + extension + 'not found')
  return blackList

def makeAll(df):
  pcmFiles = selectFiles(allFile(os.getcwd()), '.pcm')
  txtFiles = selectFiles(allFile(os.getcwd()), '.txt')
  i = 0
  pcmBlackList = []
  txtBlackList = []
  en = enDict(df)
  for row in df.itertuples():
    [pcmBlackList, change] = makeFile(pcmFiles, '.pcm', i, pcmBlackList, row)
    txtBlackList = makeTxt(txtFiles, '.txt', i, txtBlackList, row, en)
    if change:
      i+=1

def enDict(df):
  en = dict()
  for row in df.itertuples():
    fn = str(row.파일명)
    txt = str(row.영어)
    if fn not in en.keys():
      en[fn] = txt
    else:
      en[fn] += (' ' + txt)
  return en 

def makeDirs(dirs):
  for dr in dirs:
    os.makedirs(dr.replace('20200705_Flitto_Rantacar_samples','Filtto_Rantacar'))

makeDirs(allDir(os.getcwd()))
# print(toDF('.'))
df = toDF('.')
# print(enDict(df))
makeAll(df)
# print(setDF(df, '.'))

# print(selectFiles(allFile('.'), '.txt'))
