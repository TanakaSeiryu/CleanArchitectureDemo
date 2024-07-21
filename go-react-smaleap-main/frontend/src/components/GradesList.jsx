import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';


function createData(name, calories, fat, carbs, protein) {
  return { name, calories, fat, carbs, protein };
}

export default function GradesList() {
    const colmnName = ["名前", "教科", "得点"];
  const rows = [
    createData('Frozen yoghurt', 1, 6),
    createData('Ice cream sandwich', 2, 9),
    createData('Eclair', 3, 16),
    createData('Cupcake', 3, 3),
    createData('Gingerbread', 1, 16),
  ];
    return(
      <TableContainer>
        <Table aria-label="simple table">
          <TableHead>
            <TableRow>
              {colmnName.map((colmn,_) => (
                <TableCell key={_} align='left'>{colmn}</TableCell>  
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow
                key={row.name}
                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {row.name}
                </TableCell>
                <TableCell align="left" >{row.calories}</TableCell>
                <TableCell align="left" >{row.fat}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
    </TableContainer>
    )
}