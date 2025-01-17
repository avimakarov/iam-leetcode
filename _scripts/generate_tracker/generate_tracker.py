# Пробегаемся по всем каталогам с задачами:
# Для каждой директории с задачей генерируем строку типа:
#   <номер задачи>$<языки программирования>$<ссылка на leetcode>$<тег>

import os
import re
import sys

def print_info(msg: str):
    print(strings.join(" | ", ["I", msg]))

def print_warn(msg: str):
    print(strings.join(" | ", ["W", msg]))

def print_fail(msg: str) -> int:
    print(strings.join(" | ", ["X", msg]))
    return 1

def main() -> int:
    curr_folder = os.getcwd()
    if not curr_folder.endswith("iam-leetcode"):
        return print_fail("script ran not in the target folder")

if __name__ == "__main__":
    sys.exit(main())
