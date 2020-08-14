# import textdistance
import datetime
def lcs(s1, s2):
    m = [[0] * (1 + len(s2)) for i in range(1 + len(s1))]
    longest, x_longest = 0, 0
    for x in range(1, 1 + len(s1)):
        for y in range(1, 1 + len(s2)):
            if s1[x - 1] == s2[y - 1]:
                m[x][y] = m[x - 1][y - 1] + 1
                if m[x][y] > longest:
                    longest = m[x][y]
                    x_longest = x
            else:
                m[x][y] = 0
    return longest



def run(i, cp_lines):
    not_found = open('..\\testSet\\SubtTV_2017_not_found_' + str(i) + '.trn', 'rt', encoding = "UTF-8")
    lcs_match = open('..\\testSet\\SubtTV_2017_lcs_match_' + str(i) + '.trn', 'wt', encoding = "UTF-8")
    
    max_value = 0
    nf_lines = not_found.readlines()

    for nf_line in nf_lines:
        split_nf = nf_line.split(' :: ')
        file_name = split_nf[0]
        nf_txt = split_nf[1][:-1]
        lcs_list = []
        line_list = []
        for cp_line in cp_lines:
            cp_line = cp_line[:-1]
            lcs_value = lcs(nf_txt, cp_line)
            lcs_list.append(lcs_value)
            line_list.append(cp_line)
        index = lcs_list.index(max(lcs_list))
        line = line_list[index]
        lcs_match.write(nf_txt + " => " + line + "\n")

        # max_value = 0
        # for cp_line in cp_lines:
        #     cp_line = cp_line[:-1]
        #     lcs_value = lcs(nf_txt, cp_line)
        #     if lcs_value > max_value:
        #         line = cp_line
        #         max_value = lcs_value
        # print('not using list')   
        # lcs_math.write(nf_txt, "=>", line, "\n")

f = open('..\\broadcast_2017_01_03.captions_u.txt', 'rt', encoding = "UTF-8")
cp_lines = f.readlines()
run(3, cp_lines)