import os

from importStarlarkFile import importFileAs


def main():
    print("hello")




if __name__ == "__main__":
    print("Current directory:", os.getcwd())
    main()

    aasMarmeeManage = importFileAs('firstFunction', 'firstFunction.star')


    def g_extraParams():
        return aasMarmeeManage.printHello()


    print(g_extraParams())




