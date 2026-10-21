---
title: "Nametable mirroring diagrams"
url: "/category/Nametable_mirroring_diagrams/"
description: "This category collects several images used as nametable mirroring diagrams at Mirroring. They were originally designed by User:Tepples and later revised by User:rainwarrior."
wikiPage: "Category_Nametable_mirroring_diagrams"
type: category
---

This category collects several images used as nametable mirroring diagrams at [Mirroring](../../wiki/Mirroring/). They were originally designed by User:Tepples and later revised by User:rainwarrior.

These were generated with the following python3 program:

    import PIL.Image
    import PIL.ImageDraw

    tex_a = "  ****  " \
            " **  ** " \
            "**    **" \
            "**    **" \
            "********" \
            "**    **" \
            "**    **" \
            "**    **"
    tex_b = "******* " \
            "**    **" \
            "**    **" \
            "******* " \
            "**    **" \
            "**    **" \
            "**    **" \
            "******* "
    tex_c = " ****** " \
            "**    **" \
            "**    **" \
            "**      " \
            "**      " \
            "**    **" \
            "**    **" \
            " ****** "
    tex_d = "******* " \
            "**    **" \
            "**    **" \
            "**    **" \
            "**    **" \
            "**    **" \
            "**    **" \
            "******* "
    tex_2 = "        " \
            "  ***** " \
            " **   **" \
            "     ** " \
            "   ***  " \
            "  **    " \
            " **     " \
            " *******"
    tex_4 = "        " \
            "    *** " \
            "   **** " \
            "  ** ** " \
            " **  ** " \
            " *******" \
            "     ** " \
            "     ** "
    tex_0 = "        " \
            "   ***  " \
            "  *  ** " \
            " **   **" \
            " **   **" \
            " **   **" \
            "  **  * " \
            "   ***  "
    tex_Cn= "        " \
            "   **** " \
            "  **  **" \
            " **     " \
            " **     " \
            " **     " \
            "  **  **" \
            "   **** "
    tex_8 = "        " \
            "  ***** " \
            " **   **" \
            " **   **" \
            "  ***** " \
            " **   **" \
            " **   **" \
            "  ***** "
    tex = { "a":tex_a, "b":tex_b, "c":tex_c, "d":tex_d, "0":tex_0, "2":tex_2, "4":tex_4, "8":tex_8, "C":tex_Cn }

    col_a = (  0,  0,160)
    col_b = (160,  0,  0)
    col_c = (  0,120,  0)
    col_d = (120,120,  0)
    col = { "a":col_a, "b":col_b, "c":col_c, "d":col_d }

    fade = [ 130, 160, 180 ]
    fade_mirror = 160

    def fade_color(c,f):
        (r,g,b) = c
        r = f + int(((255-f)*r)/255)
        g = f + int(((255-f)*g)/255)
        b = f + int(((255-f)*b)/255)
        return (r,g,b)

    def draw_tex(img,s,col,scale,x,y):
        pixels = img.load()
        while len(s) > 0:
            t = tex[s[0]]
            s = s[1:]
            for ty in range(0,8):
                py = y + (scale * ty)
                for tx in range(0,8):
                    px = x + (scale * tx)
                    if t[0] != ' ':
                        for spy in range(0,scale):
                            for spx in range(0,scale):
                                pixels[px+spx,py+spy] = col
                    t = t[1:]
            x += (scale * 8)

    def draw_table(img,s,x,y,addr,mirr):
        pixels = img.load()
        for ty in range(0,(30*4)+1):
            for tx in range(0,(32*4)+1):
                cb = col[s]
                if (tx % (32*4) == 0) or (ty % (30*4) == 0):
                    cb = (0,0,0)
                elif (tx % 8 == 0) or (ty % 8 == 0):
                    cb = fade_color(cb,fade[0])
                elif (tx % 4 == 0) or (ty % 4 == 0):
                    cb = fade_color(cb,fade[2])
                else:
                    cb = fade_color(cb,fade[1])
                if mirr and cb != (0,0,0):
                    cb = fade_color(cb,fade_mirror)
                pixels[tx+x,ty+y] = cb
        draw_tex(img,addr,(0,0,0),1,x+8,y+8)
        draw_tex(img,s,(255,255,255),8,x+33,y+29)

    dimx = (32 * 4 * 2)+1
    dimy = (30 * 4 * 2)+1

    def mirroring_img(seq):
        subx = (32 * 4)
        suby = (30 * 4)
        img = PIL.Image.new("RGB",(dimx,dimy))
        draw_table(img,seq[0],   0,   0,"2000",False)
        draw_table(img,seq[1],subx,   0,"2400",seq[1] in seq[0:1])
        draw_table(img,seq[2],   0,suby,"2800",seq[2] in seq[0:2])
        draw_table(img,seq[3],subx,suby,"2C00",seq[3] in seq[0:3])
        img.save("mirroring_"+seq+".png")

    mirroring_img("aabb")
    mirroring_img("abab")
    mirroring_img("aaaa")
    mirroring_img("bbbb")
    mirroring_img("abcd")
    mirroring_img("abba")
    mirroring_img("abbb")
    mirroring_img("acbc")
    mirroring_img("abcc")
    mirroring_img("abbc")

    img0 = PIL.Image.open("mirroring_aaaa.png")
    img1 = PIL.Image.open("mirroring_bbbb.png")
    img = PIL.Image.new("RGB",(dimx,(dimy*2)+8),(255,255,255))
    img.paste(img0,(0,0))
    img.paste(img1,(0,dimy+8))
    img.save("mirroring_aaaa_bbbb.png")

    ox = 13
    oy = 13
    img = PIL.Image.new("RGB",(dimx+ox,dimy+oy),(255,255,255))
    img.paste(img1,(ox,0))
    img.paste(img0,(0,oy))
    img.save("mirroring_aaaa_bbbb_overlaid.png")
