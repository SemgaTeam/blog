import {
  Admin,
  Resource,
  ListGuesser,
  ShowGuesser,
  EditGuesser,
} from "react-admin";
import { Layout } from "./Layout";
import dataProvider from "./dataProvider.ts";

const App = () => (
  <Admin layout={Layout} dataProvider={dataProvider}>
    <Resource
      name="post"
      list={ListGuesser}
      show={ShowGuesser}
      edit={EditGuesser}
    />
  </Admin>
);

export default App;
