import { Edit, SimpleForm, TextInput } from "react-admin";

const PostEdit = () => (
  <Edit>
    <SimpleForm>
      <TextInput source="id" disabled />
      <TextInput source="name" />
      <TextInput source="contents" />
    </SimpleForm>
  </Edit>
);

export default PostEdit;
