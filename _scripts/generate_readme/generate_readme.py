import os
import sys

from typing import Union
from datetime import datetime

README_TEMPLATE = """
# [leetcode](https://leetcode.com/) problems solved by me

---

![UPD - {upd}](https://img.shields.io/badge/UPD-{upd}-2ea44f)
![CNT - {cnt}](https://img.shields.io/badge/CNT-{cnt}-2ea44f)

---

## Оглавление

{table_of_contents}
"""

TRACKER_GEN_FILENAME = "_scripts/_generated/tracker_gen"

class SolvedTask:
    def __init__(self, raw: str) -> None:
        tags = raw.split("$")

        self.num = tags[0]
        self.lng = tags[1]
        self.lnk = tags[2]
        self.tag = tags[3]

def print_fail(msg: str) -> int:
    print(" | ".join(["E", msg]))
    return 1

def group_by(key: str, tasks: list[SolvedTask]) -> Union[None, dict[str, list[SolvedTask]]]:
    res: dict[str, list[SolvedTask]] = dict()

    if key == "num":
        for task in tasks:
            if res.get(task.num, None) is None:
                res[task.num] = [task]
            else:
                res[task.num].append(task)

        return res

    if key == "tag":
        for task in tasks:
            if res.get(task.tag, None) is None:
                res[task.tag] = [task]
            else:
                res[task.tag].append(task)

        return res

    return None

def generate_table_of_contents(tasks: list[SolvedTask]) -> Union[int, str]:
    res = ""

    grouped_by_tag = group_by("tag", tasks)
    if grouped_by_tag is None:
        return print_fail(
            "grouping tasks by tag operation failed"
        )

    for i, k in enumerate(sorted(grouped_by_tag, key=len)):
        res += "{}. {}\n".format(i+1, k)
        for i1, v in enumerate(grouped_by_tag[k]):
            res += "    {}. [{}](https://github.com/avimakarov/iam-leetcode/tree/main/tasks/{}/code) \n".format(i1+1, v.num, v.num)

    return res

def main():
    curr_folder = os.getcwd()

    if not curr_folder.endswith("iam-leetcode"):
        return print_fail("invalid run folder")

    if not os.path.exists(TRACKER_GEN_FILENAME):
        return print_fail("base generated file not found")

    with open(TRACKER_GEN_FILENAME, "r") as file:
        tasks = [SolvedTask(line.strip()) for line in file]

    group_by_num = group_by("num", tasks)
    if group_by_num is None:
        return print_fail("grouping tasks by num operation failed")

    with open("readme.md", "w") as file:
        file.write( README_TEMPLATE.format(
            cnt = len(group_by_num),
            upd = datetime.now().strftime("%d.%m.%Y"),
            table_of_contents = generate_table_of_contents(tasks),
        ))

if __name__ == "__main__":
    sys.exit(main())
