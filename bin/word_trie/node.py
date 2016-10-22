class Node:

    def __init__(self, parent=None, letter='', is_word=False):
        self.__is_word = is_word
        self.__letter = letter
        self.__parent = parent
        self.__children = set()

    def __contains__(self, key):
        return key is self.__letter

    def set_parent(self, parent):
        self.__parent = parent

    def add_children(self, child):
        self.__children.add(child)

    def remove_children(self, child):
        self.__children.remove(child)



