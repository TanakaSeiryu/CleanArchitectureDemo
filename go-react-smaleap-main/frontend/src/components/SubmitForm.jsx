import { TextField, MenuItem, List, ListItem, Container, Button, Typography } from '@mui/material';
import { useForm } from 'react-hook-form';

import React from "react";

export default function SubmitForm() {
    const subjectBox = [
    { value: '1', label: '国語' },
    { value: '2', label: '数学' },
    { value: '3', label: '英語' },
  ];

  const gradeRules = {
    required: "得点を入力してください",
    pattern: { value: /^[0-9]+$/, message: '数字で入力してください。' }
  }
  const { register, handleSubmit,formState: { errors }} = useForm();
  const onSubmit = data => {
    console.log(data)
    };
    return (
<Container fixed sx={{
          border: 1,
          width: 1/3,
      }}>
        <List>
          <ListItem>
            <Container justifyContent="center" alignItems="center">
              <Typography variant="h3" color="text.disabled">得点記録アプリ</Typography>
            </Container>
          </ListItem>

          <ListItem >
            <TextField {...register("name", { required: true})}label="名前" fullWidth/>
          </ListItem>
          <ListItem>
            <Typography color="error.main">{errors.name && "名前を入力してください"}</Typography>
          </ListItem>

          <ListItem>
            <TextField {...register("subjectId", {required: true})} label="未選択" select fullWidth>
                {subjectBox.map((item, index) => (
                  <MenuItem key={index} value={item.value}>
                    {item.label}
                  </MenuItem>
              ))}
            </TextField>
          </ListItem>
          <ListItem><Typography color="error.main">{errors.subjectId && "科目を選択してください"}</Typography></ListItem>

          <ListItem>
            <TextField {...register("grade", gradeRules)}label="得点" fullWidth/>
          </ListItem>
          <ListItem>
            <Typography color="error.main">{errors.grade && errors.grade.message}</Typography>
          </ListItem>

          <ListItem>
            <Button variant='contained' fullWidth type="submit" onClick={handleSubmit(onSubmit)}>登録</Button>
          </ListItem>
        </List>
    </Container>
    )}