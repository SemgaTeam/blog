import {
  Datagrid,
  TextField,
  DateField,
  List,
  ReferenceField,
} from "react-admin";

export const PostList = () => (
  <List>
    <Datagrid>
      <TextField source="id" />
      <DateField source="created_at" />
      <DateField source="updated_at" />
      <TextField source="name" />
      <TextField source="contents" />
      <ReferenceField source="author_id" reference="user" sortable={false}>
        <TextField source="name" />
      </ReferenceField>
    </Datagrid>
  </List>
);

export default PostList;
