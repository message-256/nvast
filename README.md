# (N)ot (v)ery (a)bstract (s)yntax (t)ree\
fair warning i still need to test it more rigorously(it's at least in my opinion, not quite fit for use)\
i am releasing it anyway becuase the algorithm looks right and i've pretty much got like 20 minutes worth of stuff left(if all goes well)\
its basically a simpler more general (hopefully) form of an ast\
it encodes as Flat\: a string split along each recursive expression as well as the recursive expressions delimiters (eg: "1*","()")\
and as Inner\: the contents of the expression(inner has an inner to)\
