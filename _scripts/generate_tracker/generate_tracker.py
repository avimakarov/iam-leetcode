# Пробегаемся по всем каталогам с задачами:
# Для каждой директории с задачей генерируем строку типа:
#   <номер задачи>$<языки программирования>$<ссылка на leetcode>$<тег>

import os
import sys

from typing import Union

folder_item_kind_file = 1
folder_item_kind_folder = 2

class Task:
    def __init__(self, num: str) -> None:
        self.__tag = ""
        self.__lnk = ""
        self.__num = num
        self.__error = False
        self.__error_msg = ""
        self.__languages = []

        for code_file in os.listdir("{}/code".format(num)):
            if not self.__read_meta():
                self.__error = True
                self.__error_msg = "mname failed"

            if len(code_file.split(".")) != 2:
                self.__error = True
                self.__error_msg = "split failed"

            if code_file.split(".")[0] != "main":
                self.__error = True
                self.__error_msg = "fname failed"

            self.__languages.append(code_file.split(".")[1])

    def __read_meta(self) -> bool:
        with open("{}/meta".format(self.__num)) as file:
            lines = [line.rstrip() for line in file]

        if len(lines) != 2:
            return False

        if not lines[0].startswith("tag"):
            return False
        else:
            if len(lines[0].split(": ")) != 2:
                return False
            else:
                self.__tag = lines[0].split(": ")[1].strip()

        if not lines[1].startswith("lnk"):
            return False
        else:
            if len(lines[1].split(": ")) != 2:
                return False
            else:
                self.__lnk = lines[1].split(": ")[1].strip()

        return True

    def to_str(self) -> list[str]:
        return [
            "$".join([self.__num, lang, self.__lnk, self.__tag]) for lang in self.__languages
        ]

    def has_error(self) -> (bool, str):
        return self.__error, self.__error_msg

class FolderItem:
    def __init__(self, name: str, kind: int) -> None:
        self.__name: str = name
        self.__kind: int = kind

    def get_name(self) -> str:
        return self.__name

    def is_task_folder(self) -> bool:
        return (
            len(self.__name) == 4
            and
            self.__name.isdigit()
            and
            self.__kind == folder_item_kind_folder
        )

    def is_allowed(self) -> bool:
        if self.__name == ".git" and self.__kind == folder_item_kind_folder:
            return True
        if self.__name == ".idea" and self.__kind == folder_item_kind_folder:
            return True
        if self.__name == "readme.md" and self.__kind == folder_item_kind_file:
            return True
        if self.__name == "_scripts" and self.__kind == folder_item_kind_folder:
            return True
        if len(self.__name) == 4 and self.__kind == folder_item_kind_folder and self.__name.isdigit():
            return True

        return False

def print_info(msg: str):
    print(" | ".join(["I", msg]))

def print_warn(msg: str):
    print(" | ".join(["W", msg]))

def print_fail(msg: str) -> int:
    print(" | ".join(["E", msg]))
    return 1

def handle_task(num: str) -> Union[int, Task]:
    meta_path = "/".join([num, "meta"])
    code_folder = "/".join([num, "code"])

    if not os.path.exists(meta_path):
        return print_fail(
            "meta file is not exists {}".format(num)
        )

    if not os.path.exists(code_folder):
        return print_fail(
            "required folder is not exists {}".format(num)
        )

    if not len(os.listdir(code_folder)) >= 1:
        return print_fail(
            "required code files are not exists {}".format(num)
        )

    task = Task(num)
    if task.has_error()[0]:
        return print_fail(
            "something bad happened for tracking task: {} ({})".format(num, task.has_error()[1]),
        )

    return task

def folder_item_kind(name: str) -> int:
    return folder_item_kind_folder if os.path.isdir(name) else folder_item_kind_file

def main() -> int:
    tasks = []
    curr_folder = os.getcwd()
    if not curr_folder.endswith("iam-leetcode"):
        return print_fail("script ran not in the target folder")

    for item_name in os.listdir():
        item = FolderItem(
            item_name, folder_item_kind(item_name),
        )
        if item.is_allowed():
            if item.is_task_folder():
                res = handle_task(item_name)
                if res != 1:
                    tasks.append(res)
                else:
                    return print_fail(
                        "processing task error: {}".format(item_name)
                    )
        else:
            print_warn("{} is not in allowed items list".format(item.get_name()))

    with open("./_scripts/_generated/tracker_gen", "w") as f:
        for task in tasks:
            for task_str in task.to_str():
                f.write("{} \n".format(task_str))

    return 0

if __name__ == "__main__":
    sys.exit(main())
