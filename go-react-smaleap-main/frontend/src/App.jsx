import {Container} from '@mui/material';
import {GradesList} from './components/GradesList.jsx'
import SubmitForm from './components/SubmitForm.jsx'

export default function App() {
  return (
    <>
      <SubmitForm/>
      <Container ficed sx={{ width: 1 / 3, }}>
        <GradesList/>
      </Container>
      
    </>
  );
}
