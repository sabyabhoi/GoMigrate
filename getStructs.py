import os
import re

def search_files(directory, pattern):
    matches = []
    for root, _, files in os.walk(directory):
        for file in files:
            filepath = os.path.join(root, file)
            if os.path.isfile(filepath):
                with open(filepath, 'r', encoding='utf-8') as f:
                    contents = f.read()
                    for match in re.finditer(pattern, contents):
                        matches.append(match.group())
    return matches

if __name__ == "__main__":
    directory = './dao/model'
    file = open(directory + '/main.go', 'w')
    file.write('''package model

import (
  "reflect"
)

func GetStructs() []reflect.Type {
    typesArr := []reflect.Type{
    ''')
    for m in search_files(directory, r'type \w+ struct'):
        struct = m.split(' ')[1]
        file.write("\treflect.TypeOf(" + struct + "{}),\n")
    file.write('''   }
    return typesArr
}
    ''')

    file.close()
