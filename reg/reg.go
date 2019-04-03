package reg

func Known() []Plug {
  return knownPlugins[:]
}

var knownPlugins = []Plug{}

type Plug interface {
 Name() string
}

func RegisterPlugin(plug Plug) error {
  println(plug.Name())
  knownPlugins = append(knownPlugins, plug)
  return nil
}