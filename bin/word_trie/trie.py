import bin.word_trie.node as node

class Trie:

    def __init__(self):
        self.root = node.Node()
        self.number_of_words = 0

    def add_word(self, word):
        start = self.root
        for letter in word:
            if letter in start:
                pass
            else:
                start.add_children()



