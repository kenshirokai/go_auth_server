import React, { ChangeEvent, useMemo, memo } from 'react';
import styled from 'styled-components';
import { merge } from 'lodash';

/**
 *@ Elements & Styles
 */
const BaseInput = styled.input`
    font-size: 1.4rem;
    padding: 10px 15px;
    border-radius: 10px;
    outline: none;
    border:1px solid #efefef;
    width: 100%;
    box-sizing: border-box;
    transition: 0.2s ease-in;
    &:focus {
        border: 2px solid #63cdda;
    }
`

/**
 *@ Types
 */
type Options = {
    type: string
    onChange: ((e:ChangeEvent<HTMLInputElement>) => void) | undefined
}
type Props = {} & Partial<Options>
/**
 *@ default
 */
const defaultOptions = ():Options => ({
    type: 'text',
    onChange: undefined
})
/**
 *@ ReactComponent
 */
const Input: React.FC<Props> = memo((props) => {
    const options = useMemo(() => merge(defaultOptions(), {
        type: props.type,
        onChange: props.onChange
    }), [props.type, props.onChange])
    return (
        <BaseInput 
            type={options.type}
            onChange={options.onChange}
        />
    );
})

export default Input;