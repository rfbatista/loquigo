import { TextField } from '@mui/material';
import React from 'react';

type Props = {
  fullWidth?: boolean;
  className?: string;
  size?: 'small';
  placeholder?: string;
  register?: () => void;
};

const TextInput: React.FC<Props> = ({ ...props }) => {
  return <TextField {...props} />;
};

export default TextInput;
