// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example.Mod2.Inputs
{

    /// <summary>
    /// A test for namespaces (mod 2)
    /// </summary>
    public sealed class TypArgs : Pulumi.ResourceArgs
    {
        [Input("mod1")]
        public Input<Pulumi.Example.Mod1.Inputs.TypArgs>? Mod1 { get; set; }

        [Input("val")]
        public Input<string>? Val { get; set; }

        public TypArgs()
        {
            Val = "mod2";
        }
    }
}
