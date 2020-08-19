import { useEffect, RefObject } from 'react';


const useTyping = (textEl: RefObject<HTMLElement>, text: string) => {
    useEffect(() => {
        const typing = (encodeStr: string, original: string, index: number = 0) => {
            if(index == original.length){
                textEl.current!.innerText = original;
                return;
            };
            setTimeout(() => {
                textEl.current!.innerText = textEl.current!.innerText + encodeStr[index];
                typing(encodeStr, original, ++index)
            }, 40)
        }
        typing(escape(text), text)
    }, [textEl, text])
}

export default useTyping;