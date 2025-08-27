package img

import (
    "PlaceholderGen/pkg/colors"
    "bytes"
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
    "log"
    "strconv"
)

const (
    imgColorDefault  = "E5E5E5"      
    msgColorDefault  = "AAAAAA"
    imgWDefault      = 300
    imgHDefault      = 300
    fontSizeDefault  = 0
    fontfileDefault  = "wqy-zenhei.ttf"
    dpiDefault       float64 = 72
    hintingDefault   = "none"
)

// Простая генерация favicon 16×16
func GenerateFavicon() (*bytes.Buffer, error) {
    buffer := new(bytes.Buffer)
    m := image.NewRGBA(image.Rect(0, 0, 16, 16))
    clr := color.RGBA{R: 0, G: 0, B: 0, A: 255}
    draw.Draw(m, m.Bounds(), &image.Uniform{C: clr}, image.Point{0, 0}, draw.Src)

    var img image.Image = m
    if err := jpeg.Encode(buffer, img, nil); err != nil {
        return nil, err
    }
    return buffer, nil
}

// Генерация placeholder-изображения
func Generate(urlPart []string) (*bytes.Buffer, error) {
    var (
        err      error
        imgColor = imgColorDefault
        msgColor = msgColorDefault
        imgW     = imgWDefault
        imgH     = imgHDefault
        fontSize = fontSizeDefault
    )
    msg := ""

    for i, val := range urlPart {
        switch i {
        case 1:
            if val != "" {
                imgW, err = strconv.Atoi(val)
                if err != nil {
                    return nil, err
                }
            }
        case 2:
            if val != "" {
                imgH, err = strconv.Atoi(val)
                if err != nil {
                    return nil, err
                }
            }
        case 3:
            if val != "" {
                imgColor = val
            }
        case 4:
            if val != "" {
                msg = val
            }
        case 5:
            if val != "" {
                msgColor = val
            }
        case 6:
            if val != "" {
                fontSize, err = strconv.Atoi(val)
                if err != nil {
                    return nil, err
                }
            }
        }
    }

    if ((imgW > 0 || imgH > 0) && msg == "") || msg == "" {
        msg = fmt.Sprintf("%d × %d", imgW, imgH)
    }

    if fontSize == 0 {
        fontSize = imgW / 9
        if imgH < imgW {
            fontSize = imgH / 5
        }
    }

    // Парсинг цвета для фона
    hx := colors.Hex(imgColor)
    rgb, err := hx.ToRGB()
    if err != nil {
        return nil, err
    }

    m := image.NewRGBA(image.Rect(0, 0, imgW, imgH))
    imgRgba := color.RGBA{R: rgb.Red, G: rgb.Green, B: rgb.Blue, A: 255}
    draw.Draw(m, m.Bounds(), &image.Uniform{C: imgRgba}, image.Point{0, 0}, draw.Src)

    // addLabel — функция добавления надписи
    err = addLabel(m, imgW, imgH, msg, fontSize, colors.Hex(msgColor))
    if err != nil {
        log.Println("unable to add label:", err)
    }

    var img image.Image = m
    buffer := new(bytes.Buffer)
    if err := jpeg.Encode(buffer, img, nil); err != nil {
        log.Println("unable to encode image:", err)
        return nil, err
    }

    return buffer, nil
}

// Пример функции addLabel — заглушка (для ввода текста на картинку)
func addLabel(img *image.RGBA, w, h int, msg string, fontSize int, msgColor colors.Hex) error {
   
    return nil
}
