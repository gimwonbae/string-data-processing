import os

def all_dir(path):
  res = []
  for root, dirs, files in os.walk(path):
    for Dir in dirs:
      res.append(os.path.join(root, Dir))
      # res.append(Dir)
  return res

def all_file(path):
  res = []
  for root, dirs, files in os.walk(path):
    for file in files:
      res.append(os.path.join(root, file))
  return res

def change_dir_name(path):
  res = all_dir(path)
  first_dir = []
  for Dir in res:
    if len(Dir.split("\\")) == 2:
      first_dir.append(Dir)
  log = dict()
  for Dir in first_dir:
    # print(Dir.split("_"))
    # print(Dir)
    # print("_".join(Dir.split("_")[4:]))
    # print("_".join(Dir.split("_")[4:]).replace("_","-").replace("+","p"))
    # print(Dir.replace("_".join(Dir.split("_")[4:]), "_".join(Dir.split("_")[4:]).replace("_","-").replace("+","p")))
    name = "_".join(Dir.split("_")[4:])
    newName = name.replace("_","-").replace("+","p")
    if name != newName:
      log[name] = newName
      os.rename(Dir, Dir.replace(name, newName))
  secondLog = dict()

  res = all_dir(path)
  for Dir in res:
    for key in log.keys():
      if key in Dir:
        basename = Dir.split("\\")[-1]
        os.rename(Dir, Dir.replace(key, log[key]))
        secondLog[basename] = basename.replace(key, log[key])
        break
  return secondLog

def change_all_name(path, org, chg):
  dirs = all_dir(path)
  for Dir in dirs:
    if org in Dir:
      os.rename(Dir, Dir.replace(org,chg))
  files = all_file(path)
  for file in files:
    if org in file:
      os.rename(file, file.replace(org,chg))

def change_file_name(path, log):
  files = all_file(path)
  for file in files:
    for key in log.keys():
      if key in file:
        os.rename(file, file.replace(key, log[key]))
        break

# change_all_name(".","__","_")
log = change_dir_name(".")
change_file_name(".", log)
for key, value in log.items():
  print(key, '->', value)